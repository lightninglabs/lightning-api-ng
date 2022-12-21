import fs from 'fs-extra';
import path from 'path';
import {
  BUILD_CONFIG_PATH,
  BUILD_CONFIGS,
  CATEGORY_FILE,
  JSON_DIR,
  OUTPUT_DIR,
} from './constants';
import { Daemon } from './daemon';
import { templates } from './templates';
import { JsonDaemon } from './types';
import { writeCategoryJson } from './utils';

const { log, error } = console;

const main = async () => {
  const config = process.env.BUILD_CONFIG;
  if (!config) {
    log('Missing list of daemon names. Run:');
    log('BUILD_CONFIG=[lnd|labs] yarn generate-docs');
    return;
  }
  log('Generating Markdown files for Docusaurus');
  log(`Build Config: ${config}`);
  log(`Output Dir: ${path.resolve(OUTPUT_DIR)}`);

  if (fs.pathExistsSync(OUTPUT_DIR)) {
    fs.rmSync(OUTPUT_DIR, { recursive: true, force: true });
    log(`Removed output dir ${OUTPUT_DIR}`);
  }

  log('Loading templates');
  templates.load();

  log('Generating build.config.json');
  const buildConfig = BUILD_CONFIGS[config];
  if (!buildConfig) {
    log(`Error: Unable to find build config for '${config}'`);
    return;
  }
  if (process.env.PUBLIC_URL) buildConfig.baseUrl = process.env.PUBLIC_URL;
  if (process.env.BASE_PATH) buildConfig.baseUrl = process.env.BASE_PATH;
  fs.writeJsonSync(BUILD_CONFIG_PATH, buildConfig, { spaces: 2 });
  log(`Wrote ${BUILD_CONFIG_PATH}: `, buildConfig);

  const daemonNames = buildConfig.repos.map((n) => n.toLowerCase());
  log(`Scanning for JSON files in ${path.resolve(JSON_DIR)}`);
  const files = fs
    .readdirSync(JSON_DIR)
    .filter((filename) => filename.endsWith('.json'))
    .filter((filename) =>
      daemonNames.includes(path.basename(filename, '.json'))
    );
  log(`Found ${files.length} JSON files: ${files.join(', ')}`);
  if (!files.length) return;

  log('Parsing JSON files...');
  const daemons = files.map((filename) => {
    const daemonName = path.parse(filename).name;
    const jsonPath = path.join(JSON_DIR, filename);
    const json = fs.readJsonSync(jsonPath) as JsonDaemon;
    return new Daemon(daemonName, json);
  });

  log('Exporting markdown files...');
  daemons.forEach((d) => d.exportMarkdown());

  const catFilePath = path.join(OUTPUT_DIR, CATEGORY_FILE);
  writeCategoryJson(
    catFilePath,
    'API Reference',
    'Documentation for the Lightning APIs'
  );

  log('Completed generating the Markdown files');
};

try {
  main();
} catch (e) {
  error(e);
}

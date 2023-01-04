import fs from 'fs-extra';
import path from 'path';
import { CATEGORY_FILE, JSON_DIR, OUTPUT_DIR } from './constants';
import { Daemon } from './daemon';
import { templates } from './templates';
import { JsonDaemon } from './types';
import { writeCategoryJson } from './utils';

const { log, error } = console;

const main = async () => {
  log('Generating Markdown files for Docusaurus');
  log(`Output Dir: ${path.resolve(OUTPUT_DIR)}`);

  if (fs.pathExistsSync(OUTPUT_DIR)) {
    fs.rmSync(OUTPUT_DIR, { recursive: true, force: true });
    log(`Removed output dir ${OUTPUT_DIR}`);
  }

  log('Loading templates');
  templates.load();

  log(`Scanning for JSON files in ${path.resolve(JSON_DIR)}`);
  const files = fs
    .readdirSync(JSON_DIR)
    .filter((filename) => filename.endsWith('.json'));
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

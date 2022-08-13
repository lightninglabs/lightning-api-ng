import path from 'path';
import { CATEGORY_FILE, OUTPUT_DIR } from './constants';
import { Daemon } from './daemon';
import { Method } from './method';
import { JsonService } from './types';
import { snakeCase, writeCategoryJson } from './utils';

const { log } = console;

export class Service {
  name: string;
  description: string;
  protoFile: string;
  methods: Method[];

  get hasRestMethods() {
    return this.methods.some((m) => m.hasRestMethods);
  }

  get categoryInfo() {
    return {
      label: `${this.name} Service`,
      position: 0,
      link: {
        type: 'generated-index',
        description: `Documentation for the ${this.name} Service`,
      },
    };
  }

  constructor(json: JsonService, protoFile: string, daemon: Daemon) {
    log(`Creating service ${json.name} with ${json.methods.length} methods`);
    this.name = json.name;
    this.description = json.description;
    this.protoFile = protoFile;
    this.methods = json.methods.map((m) => new Method(m, daemon));
  }

  exportMarkdown(daemonName: string) {
    log(`Exporting service ${this.name}`);
    const servicePath = path.join(OUTPUT_DIR, daemonName, snakeCase(this.name));
    this.methods.forEach((m) => m.exportMarkdown(servicePath));

    const catFilePath = path.join(servicePath, CATEGORY_FILE);
    writeCategoryJson(
      catFilePath,
      `${this.name} Service`,
      `Documentation for the ${this.name} Service`
    );
  }
}

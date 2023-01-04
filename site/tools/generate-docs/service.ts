import path from 'path';
import { CATEGORY_FILE, OUTPUT_DIR } from './constants';
import { Method } from './method';
import { Package } from './package';
import { JsonService } from './types';
import { snakeCase, writeCategoryJson } from './utils';

const { log } = console;

export class Service {
  package: Package;
  fileName: string;

  name: string;
  description: string;
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

  constructor(json: JsonService, pkg: Package, fileName: string) {
    log(`Creating service ${json.name} with ${json.methods.length} methods`);
    this.package = pkg;
    this.fileName = fileName;
    this.name = json.name;
    this.description = json.description;
    this.methods = json.methods.map((m) => new Method(m, this));
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

import fs from 'fs-extra';
import path from 'path';
import { Daemon } from './daemon';
import { RestMapping } from './rest-mapping';
import { JsonMethod } from './types';
import { snakeCase } from './utils';

const { log } = console;

export class Method {
  name: string;
  description: string;
  requestFullType: string;
  requestStreaming: boolean;
  responseFullType: string;
  responseStreaming: boolean;
  restMappings: RestMapping[] = [];

  daemon: Daemon;

  get hasRestMethods() {
    return this.restMappings.length > 0;
  }

  constructor(json: JsonMethod, daemon: Daemon) {
    log(`Creating method ${json.name}`);
    this.name = json.name;
    this.description = json.description;
    this.requestFullType = json.requestFullType;
    this.requestStreaming = json.requestStreaming;
    this.responseFullType = json.responseFullType;
    this.responseStreaming = json.responseStreaming;
    this.daemon = daemon;

    if (json.restMappings) {
      this.restMappings = json.restMappings.map((m) => new RestMapping(m));
    }
  }

  exportMarkdown(servicePath: string) {
    if (!fs.pathExistsSync(servicePath)) fs.mkdirpSync(servicePath);
    const filePath = path.join(servicePath, `${snakeCase(this.name)}.md`);
    log(`Exporting ${this.name} to ${filePath}`);
    fs.writeFileSync(filePath, `# ${this.name}\n\n${this.description}`);
  }
}

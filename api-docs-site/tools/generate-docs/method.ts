import fs from 'fs-extra';
import Handlebars from 'handlebars';
import path from 'path';
import { Daemon } from './daemon';
import { RestMapping } from './rest-mapping';
import { templates } from './templates';
import { JsonMethod } from './types';
import { snakeCase } from './utils';

const { log } = console;

export class Method {
  name: string;
  description: string;
  command: string;
  requestType: string;
  requestFullType: string;
  requestStreaming: boolean;
  responseType: string;
  responseFullType: string;
  responseStreaming: boolean;
  restMappings: RestMapping[] = [];

  daemon: Daemon;

  get request() {
    return this.daemon.getMessage(this.requestFullType);
  }

  get response() {
    return this.daemon.getMessage(this.responseFullType);
  }

  get streamingDirection() {
    if (this.requestStreaming && this.responseStreaming) {
      return 'bidirectional';
    } else if (this.responseStreaming) {
      return 'server';
    } else if (this.requestStreaming) {
      return 'client';
    }
    return '';
  }

  get hasRestMethods() {
    return this.restMappings.length > 0;
  }

  constructor(json: JsonMethod, daemon: Daemon) {
    log(`Creating method ${json.name}`);
    this.name = json.name;
    this.parseDescription(json.description);
    this.requestType = json.requestType;
    this.requestFullType = json.requestFullType;
    this.requestStreaming = json.requestStreaming;
    this.responseType = json.responseType;
    this.responseFullType = json.responseFullType;
    this.responseStreaming = json.responseStreaming;
    this.daemon = daemon;

    if (json.restMappings) {
      this.restMappings = json.restMappings.map((m) => new RestMapping(m));
    }
  }

  parseDescription(description: string) {
    if (!description) return;
    const lines = description.split('\n');
    if (lines[0].includes(': `')) {
      // if the first line looks like "lncli: `closechannel`", it is a command
      this.command = lines[0];
      // log('lines[0]', lines[0]);
      this.description = lines.slice(1).join('\n');
    } else {
      this.description = description;
    }
  }

  exportMarkdown(servicePath: string) {
    if (!fs.pathExistsSync(servicePath)) fs.mkdirpSync(servicePath);

    const filePath = path.join(servicePath, `${snakeCase(this.name)}.mdx`);
    log(`Exporting ${this.name} to ${filePath}`);

    // const partials = {
    //   message: templates.message,
    //   request_message: templates.request_message,
    // };
    // const content = Mustache.render(templates.method, this, partials);
    const content = Handlebars.compile(templates.method)(this, {
      allowProtoPropertiesByDefault: true,
    });
    fs.writeFileSync(filePath, content);
  }
}

import fs from 'fs-extra';
import Handlebars from 'handlebars';
import path from 'path';
import Enum from './enum';
import { Message } from './message';
import { RestMapping } from './rest-mapping';
import { Service } from './service';
import { templates } from './templates';
import { JsonMethod } from './types';
import { snakeCase } from './utils';

const { log } = console;

export class Method {
  service: Service;

  name: string;
  description: string;
  source: string;
  commandLine: string;
  commandLineHelp: string;
  requestType: string;
  requestFullType: string;
  requestTypeSource: string;
  requestStreaming: boolean;
  responseType: string;
  responseFullType: string;
  responseTypeSource: string;
  responseStreaming: boolean;
  restMapping?: RestMapping;

  // private fields used to avoid redundant computation in getters
  private _request?: Message;
  private _response?: Message;
  private _nestedMessages?: Message[];
  private _nestedEnums?: Enum[];

  get request() {
    if (!this._request) {
      this._request = this.service.package.daemon.getMessage(
        this.requestFullType
      );
      this.restMapping?.updateMessage(this._request);
    }
    return this._request;
  }

  get response() {
    if (!this._response) {
      this._response = this.service.package.daemon.getMessage(
        this.responseFullType
      );
    }
    return this._response;
  }

  get nestedMessages() {
    if (!this._nestedMessages) {
      const messages = new Map<string, Message>();
      this.service.package.daemon.getNestedMessages(this.request, messages);
      this.service.package.daemon.getNestedMessages(this.response, messages);
      this._nestedMessages = Array.from(messages.values());
    }
    return this._nestedMessages;
  }

  get nestedEnums() {
    if (!this._nestedEnums) {
      const enums = new Map<string, Enum>();
      this.service.package.daemon.getNestedEnums(this.request, enums);
      this.service.package.daemon.getNestedEnums(this.response, enums);
      this._nestedEnums = Array.from(enums.values());
    }
    return this._nestedEnums;
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
    return !!this.restMapping;
  }

  get restMethod() {
    return this.restMapping?.method || '';
  }

  get restPath() {
    return this.restMapping?.path || '';
  }

  constructor(json: JsonMethod, service: Service) {
    log(`Parsing method ${json.name}`);
    this.service = service;
    this.name = json.name;
    this.parseDescription(json.description);
    this.source = json.source;
    this.commandLine = json.commandLine;
    this.commandLineHelp = json.commandLineHelp;
    this.requestType = json.requestType;
    this.requestFullType = json.requestFullType;
    this.requestTypeSource = json.requestTypeSource;
    this.requestStreaming = json.requestStreaming;
    this.responseType = json.responseType;
    this.responseFullType = json.responseFullType;
    this.responseTypeSource = json.responseTypeSource;
    this.responseStreaming = json.responseStreaming;

    if (json.restMappings?.length > 0) {
      this.restMapping = new RestMapping(json.restMappings[0]);
    }
  }

  parseDescription(description: string) {
    if (!description) return;
    const lines = description.split('\n');
    if (lines[0].includes(': `')) {
      // if the first line looks like "lncli: `closechannel`", it is
      // a command, so skip it
      this.description = lines.slice(1).join('\n');
    } else {
      this.description = description;
    }
  }

  exportMarkdown(servicePath: string) {
    if (!fs.pathExistsSync(servicePath)) fs.mkdirpSync(servicePath);

    const filePath = path.join(servicePath, `${snakeCase(this.name)}.mdx`);
    log(`Exporting ${this.name} to ${filePath}`);

    const content = Handlebars.compile(templates.method)(this, {
      allowProtoPropertiesByDefault: true,
    });
    fs.writeFileSync(filePath, content);
  }
}

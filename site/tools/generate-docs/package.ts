import { Daemon } from './daemon';
import Enum from './enum';
import { Message } from './message';
import { Service } from './service';
import { JsonProtoFile } from './types';

const { log } = console;

export class Package {
  daemon: Daemon;

  name: string;
  description: string;
  messages = new Map<string, Message>();
  enums = new Map<string, Enum>();
  services: Service[] = [];
  experimental = false;

  constructor(name: string, daemon: Daemon) {
    log(`Creating package ${name}`);
    this.daemon = daemon;
    this.name = name;
  }

  addProtoFile(json: JsonProtoFile) {
    log(`Adding proto file ${json.name} to package ${json.package}`);
    this.name = json.package;
    this.description = json.description;

    log(`Adding ${json.messages.length} messages in ${json.package}`);
    json.messages.forEach((m) =>
      this.messages.set(m.longName, new Message(m, this, json.name))
    );

    log(`Adding ${json.enums.length} enums in ${json.package}`);
    json.enums.forEach((e) =>
      this.enums.set(e.longName, new Enum(e, this, json.name))
    );

    log(`Adding ${json.services.length} services in ${json.package}`);
    json.services.forEach((s) =>
      this.services.push(new Service(s, this, json.name))
    );
  }

  exportMarkdown(daemonName: string) {
    log(`Exporting package ${this.name}`);
    this.services.forEach((s) => s.exportMarkdown(daemonName));
  }
}

import Field from './field';
import { Package } from './package';
import { JsonMessage } from './types';

const { log } = console;

export class Message {
  package: Package;

  name: string;
  longName: string;
  fullName: string;
  description: string;
  fields: Field[];

  get hasFields() {
    return this.fields.length > 0;
  }

  constructor(json: JsonMessage, pkg: Package) {
    log(`Creating message ${json.name} with ${json.fields.length} fields`);
    this.package = pkg;
    this.name = json.name;
    this.longName = json.longName;
    this.fullName = json.fullName;
    this.description = json.description;
    this.fields = json.fields.map((f) => new Field(f, this));
  }
}

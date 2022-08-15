import Field from './field';
import { JsonMessage } from './types';

const { log } = console;

export class Message {
  name: string;
  description: string;
  protoFile: string;
  fields: Field[];

  get hasFields() {
    return this.fields.length > 0;
  }

  constructor(json: JsonMessage, protoFile: string) {
    log(`Creating message ${json.name} with ${json.fields.length} fields`);
    this.name = json.name;
    this.description = json.description;
    this.protoFile = protoFile;
    this.fields = json.fields.map((f) => new Field(f));
  }
}

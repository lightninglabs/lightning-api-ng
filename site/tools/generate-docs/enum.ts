import { Package } from './package';
import { JsonEnum } from './types';

const { log } = console;

export default class Enum {
  package: Package;
  fileName: string;

  name: string;
  longName: string;
  fullName: string;
  description: string;
  values: {
    name: string;
    number: string;
    description: string;
  }[];

  constructor(json: JsonEnum, pkg: Package, fileName: string) {
    log(`Creating enum ${json.name} with ${json.values.length} values`);
    this.package = pkg;
    this.fileName = fileName;
    this.name = json.name;
    this.longName = json.longName;
    this.fullName = json.fullName;
    this.description = json.description
      .replace(/\n/g, ' ') // replace newlines with spaces
      .replace(/\|/g, '\\|'); // escape pipe "|" chars
    this.values = json.values.map((v) => ({
      name: v.name,
      number: v.number,
      description: v.description
        .replace(/\n/g, ' ') // replace newlines with spaces
        .replace(/\|/g, '\\|'), // escape pipe "|" chars
    }));
  }
}

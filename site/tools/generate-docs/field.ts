import { JsonField } from './types';

export default class Field {
  name: string;
  description: string;
  label: string;
  type: string;
  fullType: string;
  isoneof: boolean;

  get grpcType() {
    if (this.label === 'repeated') return `array ${this.type}`;
    return this.type;
  }

  constructor(json: JsonField) {
    this.name = json.name;
    this.description = json.description
      .replace(/\n/g, ' ') // replace newlines with spaces
      .replace(/\|/g, '\\|'); // escape pipe "|" chars
    this.label = json.label;
    this.type = json.type;
    this.fullType = json.fullType;
    this.isoneof = json.isoneof;
  }
}

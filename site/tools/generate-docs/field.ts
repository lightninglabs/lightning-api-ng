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

  get linkedType() {
    if (this.fullType.includes('.')) {
      const hash = this.fullType.toLowerCase().split('.').join('');
      return `[\`${this.grpcType}\`](#${hash})`;
    } else {
      return `\`${this.grpcType}\``;
    }
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

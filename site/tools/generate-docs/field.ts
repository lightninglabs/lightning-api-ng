import { JsonField } from './types';

export default class Field {
  name: string;
  description: string;
  label: string;
  type: string;
  fullType: string;
  isoneof: boolean;

  /** This value is set externally */
  restType = 'unknown';
  restPlacement = 'unknown';

  get grpcType() {
    if (this.label === 'repeated') return `${this.type}[]`;
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

  get encodingTip() {
    if (this.type === 'bytes') {
      if (this.restPlacement === 'query') {
        return '<Tip>Use base64 & URL encoding.<br /> See [REST Encoding](../#rest-encoding).</Tip>';
      }
      if (this.restPlacement === 'body') {
        return '<Tip>Use base64 encoding.<br /> See [REST Encoding](../#rest-encoding).</Tip>';
      }
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

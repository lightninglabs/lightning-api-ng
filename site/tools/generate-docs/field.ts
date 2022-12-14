import { Message } from './message';
import { JsonField } from './types';

export default class Field {
  message: Message;

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
    let encoding: string;
    if (this.type === 'bytes') {
      if (this.restPlacement === 'query') encoding = 'base64 & URL';
      if (this.restPlacement === 'body') encoding = 'base64';
      if (encoding) {
        return [
          '<Tip>',
          `Use ${encoding} encoding in the ${this.restPlacement}.`,
          '<br />',
          `See [REST Encoding](/docs/api/${this.message.package.daemon.name}/#rest-encoding).`,
          '</Tip>',
        ].join('');
      }
    }
  }

  constructor(json: JsonField, message: Message) {
    this.message = message;
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

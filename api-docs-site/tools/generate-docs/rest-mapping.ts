import { JsonRestMapping } from './types';

export class RestMapping {
  method: string;
  path: string;

  constructor(json: JsonRestMapping) {
    this.method = json.method;
    this.path = json.path;
  }
}

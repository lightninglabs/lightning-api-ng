import Field from './field';
import { Message } from './message';
import { JsonRestMapping } from './types';

export interface RestParameter {
  name: string;
  format: string;
  type: string;
  in: string;
}

export class RestMapping {
  method: string;
  path: string;
  /**
   * Maps from parameter name to its info. Keys could be:
   * 'sat_per_byte' or 'channel_point.funding_txid_str'
   */
  parameters = new Map<string, RestParameter>();
  /**
   * Indicates if there is a generic `body` parameter which
   * holds the full request message
   */
  hasBodyParams = false;

  constructor(json: JsonRestMapping) {
    this.method = json.method;
    this.path = json.path;

    json.details?.parameters?.forEach((p) => {
      if (this.method === 'POST' && p.name === 'body' && p.schema?.$ref) {
        // all of the parameters are in the body
        this.hasBodyParams = true;
      } else {
        this.parameters.set(p.name, p);
      }
    });
  }

  updateMessage(msg: Message) {
    // update REST type and placement from restMappings
    msg.fields.forEach((field) => {
      // skip fields that have already been set from the restTypes parsing in `rest-types.ts`
      if (field.restType !== 'unknown' && field.restPlacement !== 'unknown')
        return;

      if (field.fullType.includes('.')) {
        this.updateStructField(field);
      } else {
        this.updateField(field);
      }

      // if no parameter was found to update the placement, this is because
      // the restMapping only has a `body` parameter for POST requests
      if (this.hasBodyParams && field.restPlacement === 'unknown') {
        field.restPlacement = 'body';
      }
    });
  }

  updateField(field: Field) {
    const param = this.parameters.get(field.name);
    if (param) {
      field.restType = param.type;
      field.restPlacement = param.in;
    }
  }

  /**
   * Updates the field when there are multiple parameters
   */
  updateStructField(field: Field) {
    const params = this.getParamsByPrefix(field.name);
    // the REST type may have been set
    if (field.restType === 'unknown') field.restType = 'object';
    field.restPlacement = this.getCombinedPlacement(params);
    // update the field's description to include details about the placement of each param
    const details = params
      .map((p) => `\`${p.name}\`: \`${p.type}\` in \`${p.in}\``)
      .join('<br />');
    if (details)
      field.description += `<h4>Nested REST Parameters</h4><p>${details}</p>`;
  }

  /**
   * Returns a list of parameters whose name starts with a prefix
   */
  getParamsByPrefix(prefix: string) {
    return Array.from(this.parameters.values()).filter((p) =>
      p.name.startsWith(prefix)
    );
  }

  /**
   * Returns the placement of all parameters. If they are the same, return the placement
   * otherwise return 'mixed'
   */
  getCombinedPlacement(params: RestParameter[]) {
    if (params.length === 0) return 'unknown';
    // create an array of unique placements
    const placements = params
      .map((p) => p.in)
      .filter((pIn, all) => pIn !== all[0]);
    return placements.length > 1 ? 'mixed' : placements[0];
  }
}

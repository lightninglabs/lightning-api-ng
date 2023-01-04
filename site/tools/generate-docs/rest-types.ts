import { Message } from './message';
import { JsonRestEnum, JsonRestObject, JsonRestTypes } from './types';

const { log } = console;

export class RestTypes {
  /**
   * A mapping from a message's full field name to the REST type
   * ex: frdrpcCloseRecommendationsResponse.considered_channels -> integer
   */
  fieldTypes = new Map<string, string>();
  /**
   * A mapping from an enum's name to the REST type
   * ex: ForceClosedChannelAnchorState -> string
   */
  enumTypes = new Map<string, string>();

  constructor(types: JsonRestTypes) {
    const entries = Object.entries(types);
    log(`Creating ${entries.length} RestTypes`);
    entries
      .filter(([name, schema]) => schema.type !== 'object')
      .forEach(([name, schema]) =>
        this.parseEnum(name, schema as JsonRestEnum)
      );
    entries
      .filter(([name, schema]) => schema.type === 'object')
      .forEach(([name, schema]) =>
        this.parseSchema(name, schema as JsonRestObject)
      );
  }

  parseEnum(typeName: string, schema: JsonRestEnum) {
    this.enumTypes.set(typeName, schema.type);
  }

  parseSchema(typeName: string, schema: JsonRestObject) {
    // skip if there are no properties defined
    if (!schema.properties) return;
    Object.entries(schema.properties).forEach(([fieldName, details]) => {
      const fullName = `${typeName}.${fieldName}`;
      let fieldType = 'unknown';
      if (details.$ref) {
        // $ref looks like: '#/definitions/InvoiceInvoiceState'
        // check if it is an enum or fallback to 'object'
        const refName = details.$ref.substring(
          details.$ref.lastIndexOf('/') + 1
        );
        const enumType = this.enumTypes.get(refName);
        fieldType = enumType || 'object';
      } else if (details.type) {
        fieldType = details.type;
      } else {
        throw new Error(`Could not determine the field type for ${fullName}`);
      }
      this.fieldTypes.set(fullName, fieldType);
    });
  }

  updateMessage(message: Message) {
    const msgRestType = this.convertGrpcTypeToRest(message.fullName);
    message.fields.forEach((field) => {
      const restType = `${msgRestType}.${field.name}`;
      if (this.fieldTypes.has(restType)) {
        field.restType = this.fieldTypes.get(restType);
      }
    });
  }

  /**
   * Converts the gRPC full type to a REST type
   * Examples:
   *   lnrpc.AddInvoiceResponse -> lnrpcAddInvoiceResponse
   */
  convertGrpcTypeToRest(typeName: string) {
    const parts = typeName.split('.');
    return `${parts[0]}${parts[parts.length - 1]}`;
  }
}

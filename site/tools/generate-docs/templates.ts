import fs from 'fs';
import Handlebars from 'handlebars';
import path from 'path';

const TEMPLATES_DIR = path.join('tools', 'templates');

const { log } = console;

class Templates {
  /**
   * The template for RPC methods
   */
  method: string;
  /**
   * The template partial for RPC messages
   */
  message: string;
  /**
   * The template partial to use for request messages. This differs from
   * the 'message' template in that is contains a "REST Placement" column
   * in the table of fields
   */
  request_message: string;
  /**
   * The template for enums
   */
  enumeration: string;

  loadTemplate(fileName: string) {
    const filePath = path.join(TEMPLATES_DIR, `${fileName}.md`);
    log(`Loading template ${filePath}`);
    return fs.readFileSync(filePath).toString();
  }

  load() {
    this.method = this.loadTemplate('method');
    this.message = this.loadTemplate('message');
    this.request_message = this.loadTemplate('request_message');
    this.enumeration = this.loadTemplate('enum');

    Handlebars.registerPartial('partial_message', this.message);
    Handlebars.registerPartial('partial_request_message', this.request_message);
    Handlebars.registerPartial('partial_enum', this.enumeration);
  }
}

export const templates = new Templates();

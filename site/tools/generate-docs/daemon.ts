import fs from 'fs-extra';
import path from 'path';
import { OUTPUT_DIR } from './constants';
import Enum from './enum';
import { Message } from './message';
import { Package } from './package';
import { RestTypes } from './rest-types';
import { JsonDaemon } from './types';

const { log } = console;

export class Daemon {
  name: string;
  packages = new Map<string, Package>();
  restTypes: RestTypes;

  constructor(daemonName: string, json: JsonDaemon) {
    log(`Creating daemon ${daemonName} with ${json.files.length} proto files`);
    this.name = daemonName;

    this.restTypes = new RestTypes(json.restTypes);

    json.files.forEach((f) => {
      let pkg = this.packages.get(f.package);
      if (!pkg) {
        pkg = new Package(f.package);
        this.packages.set(f.package, pkg);
      }
      pkg.addProtoFile(f, this);
    });
  }

  get camelName() {
    if (this.name === 'lnd') return 'LND';
    return this.name[0].toUpperCase() + this.name.substring(1);
  }

  getMessage(fullType: string, throwError = true) {
    // split "lnrpc.Invoice.InvoiceState" into "lnrpc" and "Invoice.InvoiceState"
    const period = fullType.indexOf('.');
    const pkgName = fullType.substring(0, period);
    const msgType = fullType.substring(period + 1);

    if (!this.packages.has(pkgName)) {
      if (throwError) {
        throw new Error(`Cannot find package ${pkgName} for ${fullType}`);
      } else {
        return;
      }
    }
    const pkg = this.packages.get(pkgName);

    if (!pkg.messages.has(msgType)) {
      if (throwError) {
        throw new Error(
          `Cannot find message ${msgType} for ${fullType} in the ${pkgName} package`
        );
      } else {
        return;
      }
    }
    const msg = pkg.messages.get(msgType);
    this.restTypes.updateMessage(msg);
    return msg;
  }

  getNestedMessages(message: Message, allMessages: Map<string, Message>) {
    message.fields
      .map((f) => f.fullType)
      // only include the non-native field types (ex: lnrpc.OutPoint)
      .filter((t) => t.includes('.'))
      // add the messages for each type
      .forEach((t) => {
        const msg = this.getMessage(t, false);
        // add the message to the map if it's not in there already
        if (msg) {
          // add the message to the map
          allMessages.set(t, msg);

          // add the nested messages for this message
          this.getNestedMessages(msg, allMessages);
        }
      });
  }

  getEnum(fullType: string, throwError = true) {
    // split "lnrpc.Invoice.InvoiceState" into "lnrpc" and "Invoice.InvoiceState"
    const period = fullType.indexOf('.');
    const pkgName = fullType.substring(0, period);
    const enumType = fullType.substring(period + 1);

    if (!this.packages.has(pkgName)) {
      if (throwError) {
        throw new Error(`Cannot find package ${pkgName} for ${fullType}`);
      } else {
        return;
      }
    }
    const pkg = this.packages.get(pkgName);

    if (!pkg.enums.has(enumType)) {
      if (throwError) {
        throw new Error(
          `Cannot find enum ${enumType} for ${fullType} in the ${pkgName} package`
        );
      } else {
        return;
      }
    }
    return pkg.enums.get(enumType);
  }

  getNestedEnums(message: Message, allEnums: Map<string, Enum>) {
    message.fields
      .map((f) => f.fullType)
      // only include the non-native field types (ex: lnrpc.OutPoint)
      .filter((t) => t.includes('.'))
      // add the messages for each type
      .forEach((t) => {
        const enu = this.getEnum(t, false);
        if (enu) {
          allEnums.set(t, enu);
          return;
        }
        // if the enum wasn't found, look for a nested message which
        // may have enum fields
        const msg = this.getMessage(t, false);
        if (msg) {
          // search the nested messages for more enums
          this.getNestedEnums(msg, allEnums);
        }
      });
  }

  exportMarkdown() {
    log(`Exporting daemon ${this.camelName}`);

    this.packages.forEach((f) => f.exportMarkdown(this.name));

    const indexFilePath = path.join(OUTPUT_DIR, this.name, 'index.md');
    fs.writeFileSync(indexFilePath, `# ${this.camelName} API`);
  }
}

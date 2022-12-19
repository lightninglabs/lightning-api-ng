import { Method } from './method';

/**
 * This class represents the object that is used to render the markdown
 * for the code samples for each method. Using this class reduces the
 * amount of logic needed in the template files
 */
export default class CodeSamples {
  method: Method;

  constructor(method: Method) {
    this.method = method;
  }

  get daemonName() {
    return this.method.service.package.daemon.name;
  }

  get packageName() {
    return this.method.service.package.name;
  }

  get serviceName() {
    return this.method.service.name;
  }

  get loaderFiles() {
    const pkg = this.method.service.package;
    if (pkg.daemon.name === 'lnd') {
      return pkg.fileName === 'lightning.proto'
        ? `'lightning.proto'`
        : `['lightning.proto', '${pkg.fileName}']`;
    } else {
      return `'${pkg.fileName}'`;
    }
  }

  get macaroonPath() {
    return this.daemonName === 'lnd'
      ? `LND_DIR/data/chain/bitcoin/regtest/admin.macaroon`
      : `${this.daemonName.toUpperCase()}_DIR/regtest/${
          this.daemonName
        }.macaroon`;
  }

  get requiresMacaroon() {
    const anonServices = ['lnrpc.WalletUnlocker', 'lnrpc.State'];
    return !anonServices.includes(`${this.packageName}.${this.serviceName}`);
  }

  get grpcPort() {
    return this.method.service.package.daemon.grpcPort;
  }

  get restPort() {
    return this.method.service.package.daemon.restPort;
  }

  get requestFields() {
    return this.method.request.fields;
  }

  get responseFields() {
    return this.method.response.fields;
  }

  get isUnary() {
    return this.method.streamingDirection === '';
  }

  get isServerStreaming() {
    return this.method.streamingDirection === 'server';
  }

  get isClientStreaming() {
    return this.method.streamingDirection === 'client';
  }

  get isBidirectionalStreaming() {
    return this.method.streamingDirection === 'bidirectional';
  }

  get isStreaming() {
    return !this.isUnary;
  }

  get isRestPost() {
    return this.method.restMethod === 'POST';
  }
}

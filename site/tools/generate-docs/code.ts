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

  get macaroonName() {
    return this.daemonName === 'lnd' ? 'admin' : this.daemonName;
  }

  get requiresMacaroon() {
    return this.serviceName !== 'WalletUnlocker';
  }

  get grpcPort() {
    return this.method.service.package.daemon.grpcPort;
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
}

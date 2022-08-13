export interface JsonRestMapping {
  method: string;
  path: string;
}

export interface JsonMethod {
  name: string;
  description: string;
  requestType: string;
  requestFullType: string;
  requestStreaming: boolean;
  responseType: string;
  responseFullType: string;
  responseStreaming: boolean;
  restMappings: JsonRestMapping[];
}

export interface JsonField {
  name: string;
  description: string;
  label: string;
  type: string;
  longType: string;
  fullType: string;
  ismap: boolean;
  isoneof: boolean;
  defaultValue: string;
}

export interface JsonMessage {
  name: string;
  description: string;
  fields: JsonField[];
}

export interface JsonService {
  name: string;
  description: string;
  methods: JsonMethod[];
}

export interface JsonProtoFile {
  name: string;
  description: string;
  package: string;
  messages: JsonMessage[];
  services: JsonService[];
}

export interface JsonDaemon {
  files: JsonProtoFile[];
}

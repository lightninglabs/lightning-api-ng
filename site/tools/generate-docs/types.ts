export interface JsonRestMapping {
  method: string;
  path: string;
}

export interface JsonMethod {
  name: string;
  description: string;
  source: string;
  commandLine: string;
  commandLineHelp: string;
  requestType: string;
  requestFullType: string;
  requestTypeSource: string;
  requestStreaming: boolean;
  responseType: string;
  responseFullType: string;
  responseTypeSource: string;
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
  longName: string;
  fullName: string;
  description: string;
  fields: JsonField[];
}

export interface JsonEnum {
  name: string;
  longName: string;
  fullName: string;
  description: string;
  values: {
    name: string;
    number: string;
    description: string;
  }[];
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
  enums: JsonEnum[];
  services: JsonService[];
}

export interface JsonDaemon {
  files: JsonProtoFile[];
}

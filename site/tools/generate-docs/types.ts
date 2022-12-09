export interface JsonRestEnum {
  default: string;
  enum: string[];
  type: string;
}

export interface JsonRestField {
  $ref?: string;
  type?: string;
  description: string;
  title?: string;
  format?: string;
  items?: {
    $ref?: string;
    type?: string;
  };
}

export interface JsonRestObject {
  type: string;
  properties: { [key: string]: JsonRestField };
}

export type JsonRestSchemaObject = JsonRestObject | JsonRestEnum;

export interface JsonRestTypes {
  [key: string]: JsonRestSchemaObject;
}

export interface JsonRestParameter {
  name: string;
  format: string;
  type: string;
  in: string;
  schema?: {
    $ref: string;
  };
}

export interface JsonRestMapping {
  method: string;
  path: string;
  details: {
    operationId: string;
    parameters: JsonRestParameter[];
  };
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
  restTypes: JsonRestTypes;
}

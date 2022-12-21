import path from 'path';
import { BuildConfig } from './types';

export const JSON_DIR = path.join('..', 'build');
export const OUTPUT_DIR = path.join('docs', 'api');
export const CATEGORY_FILE = '_category_.json';

export const BUILD_CONFIG_PATH = 'build.config.json';
export const BUILD_CONFIGS: Record<string, BuildConfig> = {
  lnd: {
    title: 'LND API Reference',
    url: 'https://api.lightning.comunnity',
    baseUrl: '/',
    repos: ['LND'],
  },
  labs: {
    title: 'Lightning Labs API Reference',
    url: 'https://lightning.engineering',
    baseUrl: '/api-docs/',
    repos: ['Loop', 'Pool', 'Faraday', 'Taro'],
  },
};

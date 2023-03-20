// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'Lightning Labs API Reference',
  tagline: 'Dinosaurs are cool',
  url: 'https://lightning.engineering',
  baseUrl: '/api-docs/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon-32x32.png',

  // Even if you don't use internalization, you can use this field to set useful
  // metadata like html lang. For example, if your site is Chinese, you may want
  // to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  themes: [
    // ... Your other themes.
    [
      require.resolve('@easyops-cn/docusaurus-search-local'),
      /** @type {import("@easyops-cn/docusaurus-search-local").PluginOptions} */
      ({
        // ... Your options.
        // `hashed` is recommended as long-term-cache of index file is possible.
        hashed: true,
        // Base route path(s) of docs. Slash at beginning is not required.
        // Note: for docs-only mode, this needs to be the same as routeBasePath
        // in your @docusaurus/preset-classic config e.g., "/".
        docsRouteBasePath: '/',
        // Whether to index blog.
        indexBlog: false,
      }),
    ],
  ],

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          routeBasePath: '/',
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            'https://github.com/lightninglabs/lightning-api-ng/edit/main/',
        },
        blog: false,
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      navbar: {
        title: 'Lightning Labs API Reference',
        logo: {
          alt: 'Lightning Labs Logo',
          src: 'img/icon-48x48.png',
        },
        items: [
          {
            type: 'doc',
            docId: 'api/lnd/index',
            position: 'left',
            label: 'LND',
          },
          {
            type: 'doc',
            docId: 'api/loop/index',
            position: 'left',
            label: 'Loop',
          },
          {
            type: 'doc',
            docId: 'api/pool/index',
            position: 'left',
            label: 'Pool',
          },
          {
            type: 'doc',
            docId: 'api/faraday/index',
            position: 'left',
            label: 'Faraday',
          },
          {
            type: 'doc',
            docId: 'api/taro/index',
            position: 'left',
            label: 'Taro',
          },
          {
            type: 'doc',
            docId: 'api/lit/index',
            position: 'left',
            label: 'LiT',
          },
          {
            href: 'https://github.com/lightninglabs/lightning-api-ng/issues',
            label: 'Feedback',
            position: 'right',
          },
        ],
      },
      footer: {
        style: 'dark',
        links: [
          {
            title: 'Docs',
            items: [
              {
                label: 'LND',
                to: 'api/lnd',
              },
              {
                label: 'Loop',
                to: 'api/loop',
              },
              {
                label: 'Pool',
                to: 'api/pool',
              },
              {
                label: 'Faraday',
                to: 'api/faraday',
              },
              {
                label: 'Taro',
                to: 'api/taro',
              },
              {
                label: 'LiT',
                to: 'api/lit',
              },
            ],
          },
          {
            title: 'Community',
            items: [
              {
                label: "Builder's Guide",
                href: 'https://docs.lightning.engineering',
              },
              {
                label: 'Slack',
                href: 'https://lightning.engineering/slack.html',
              },
              {
                label: 'Twitter',
                href: 'https://twitter.com/lightning',
              },
            ],
          },
          {
            title: 'Github',
            items: [
              {
                label: 'lightningnetwork/lnd',
                href: 'http://github.com/lightningnetwork/lnd',
              },
              {
                label: 'lightninglabs/loop',
                href: 'http://github.com/lightninglabs/loop',
              },
              {
                label: 'lightninglabs/pool',
                href: 'http://github.com/lightninglabs/pool',
              },
              {
                label: 'lightninglabs/faraday',
                href: 'http://github.com/lightninglabs/faraday',
              },
              {
                label: 'lightninglabs/taro',
                href: 'http://github.com/lightninglabs/taro',
              },
              {
                label: 'lightninglabs/lit',
                href: 'http://github.com/lightninglabs/lit',
              },
            ],
          },
        ],
        copyright: `Copyright © ${new Date().getFullYear()} Lightning Labs, Inc.`,
      },
      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
      },
    }),
};

module.exports = config;

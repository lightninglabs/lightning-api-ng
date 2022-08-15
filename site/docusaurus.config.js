// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'Lightning API Docs',
  tagline: 'Dinosaurs are cool',
  url: 'https://api-docs.lightning.engineering/',
  baseUrl: '/api-docs-site/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.ico',

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: 'lightninglabs', // Usually your GitHub org/user name.
  projectName: 'api-docs-site', // Usually your repo name.

  // Even if you don't use internalization, you can use this field to set useful
  // metadata like html lang. For example, if your site is Chinese, you may want
  // to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl: 'https://github.com/lightninglabs/api-docs-site/edit/main/',
        },
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
        title: 'Lightning API Docs',
        logo: {
          alt: 'Lightning API Docs Logo',
          src: 'img/logo.svg',
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
            href: 'https://github.com/lightninglabs/api-docs-site',
            label: 'GitHub',
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
                label: 'Lightning API Docs',
                to: '/docs/lnd',
              },
            ],
          },
          {
            title: 'Community',
            items: [
              {
                label: 'Stack Overflow',
                href: 'https://stackoverflow.com/questions/tagged/lightning',
              },
              {
                label: 'Slack',
                href: 'https://lightning.engineering/slack.html',
              },
              {
                label: 'Twitter',
                href: 'https://twitter.com/api-docs-site',
              },
            ],
          },
          {
            title: 'More',
            items: [
              {
                label: 'Blog',
                to: 'https://lightning.engineering/blog',
              },
              {
                label: 'GitHub',
                href: 'https://github.com/lightninglabs/api-docs-site',
              },
            ],
          },
        ],
        copyright: `Copyright © ${new Date().getFullYear()} Lightning Labs, Inc. Built with Docusaurus.`,
      },
      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
      },
    }),
};

module.exports = config;

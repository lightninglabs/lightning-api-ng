// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');
const buildConfig = require('./build.config.json');

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: buildConfig.title,
  tagline: 'Dinosaurs are cool',
  url: buildConfig.url,
  baseUrl: buildConfig.baseUrl,
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
        title: buildConfig.title,
        logo: {
          alt: 'Lightning Labs Logo',
          src: 'img/icon-48x48.png',
        },
        items: [
          ...buildConfig.repos.map((label) => ({
            type: 'doc',
            docId: `api/${label.toLowerCase()}/index`,
            position: 'left',
            label,
          })),
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
              ...buildConfig.repos.map((label) => ({
                label,
                to: `api/${label.toLowerCase()}`,
              })),
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
            items: buildConfig.repos.map((label) => {
              const repoName = label.toLowerCase();
              const orgName =
                repoName === 'lnd' ? 'lightningnetwork' : 'lightninglabs';
              return {
                label: `${orgName}/${repoName}`,
                href: `https://github.com/${orgName}/${repoName}`,
              };
            }),
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

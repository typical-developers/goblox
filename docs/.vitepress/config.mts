import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  cleanUrls: true,

  title: "Goblox Documentation",
  description: "Documentation for the Goblox Go library.",

  head: [
    ['link', { rel: 'icon', href: '/goblox_icon.png' }],

    ['meta', { property: 'og:type', content: 'website' }],
    ['meta', { property: 'og:site_name', content: 'Goblox' }],
    ['meta', { property: 'og:title', content: 'Goblox Documentation' }],
    ['meta', { property: 'og:description', content: "Easily integrate Roblox's Open Cloud & legacy APIs into your Go projects." }],
    ['meta', { property: 'og:url', content: "https://goblox.dev" }],
    ['meta', { property: 'og:image', content: `https://goblox.dev/goblox_icon.png` }],

    ['meta', { name: 'twitter:card', content: 'summary' }],

    ['script', { id: "discord:component-embed", type: "application/json" }, JSON.stringify({
      component: {
        type: 17,
        accent_color: 0x6ad7e5,
        components: [
          {
            type: 9,
            components: [
              { type: 10, content: "## Goblox Documentation\nEasily integrate Roblox's Open Cloud & legacy APIs into your Go projects." },
            ],
            accessory: { type: 11, media: { url: `https://goblox.dev/goblox_icon.png` } },
          },
          {
            type: 1,
            components: [
              { type: 2, style: 5, label: 'OpenCloud', url: `https://goblox.dev/guides/opencloud/authentication` },
              { type: 2, style: 5, label: 'Legacy', url: `https://goblox.dev/guides/legacy/authentication` },
              { type: 2, style: 5, label: 'GitHub', url: 'https://github.com/typical-developers/goblox' },
            ],
          },
        ],
      },
    })]
  ],

  markdown: {
    lineNumbers: true
  },

  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    siteTitle: "Goblox",
    logo: '/goblox_icon.png',
    lastUpdated: {},
    editLink: {
      pattern: "https://github.com/typical-developers/goblox/edit/main/docs/:path"
    },

    search: {
      provider: 'local',
    },
    outline: [2, 4],

    nav: [
      { text: 'Home', link: '/' },
      { text: 'Introduction', link: '/getting-started/welcome' },
      { text: 'OpenCloud', link: '/guides/opencloud/authentication' },
      { text: 'Legacy', link: '/guides/legacy/authentication' },
    ],

    sidebar: [
      {
        text: 'Introduction',
        items: [
          { text: 'About Goblox', link: '/getting-started/welcome' },
        ],
      },
      {
        text: 'OpenCloud',
        collapsed: false,
        items: [
          { text: 'Authentication', link: '/guides/opencloud/authentication' },
          { text: 'Pagination', link: '/guides/opencloud/pagination' },
          { text: 'Luau Execution', link: '/guides/opencloud/luau-execution' },
          { text: 'Polling Endpoints', link: '/guides/opencloud/polling-endpoints' },
        ],
      },
      {
        text: 'Legacy',
        collapsed: false,
        items: [
          { text: 'Authentication', link: '/guides/legacy/authentication' },
        ],
      },
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/typical-developers/goblox' },
      { icon: 'discord', link: 'https://discord.gg/typical' },
      { icon: 'twitter', link: 'https://twitter.com/typicaldevelops' },
      { icon: 'roblox', link: 'https://www.roblox.com/communities/2700233/Typical-Developers' },
    ]
  }
})

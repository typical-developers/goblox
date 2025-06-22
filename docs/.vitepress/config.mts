import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Goblox Documentation",
  description: "Documentation for the Goblox Go library.",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Introduction', link: '/getting-started/welcome' },
      { text: 'Guides', link: '/guides/opencloud/client' },
      { text: 'Packages', link: '/getting-started/packages/methodutil' },
    ],

    sidebar: [
      {
        text: 'Introduction',
        items: [
          { text: 'About Goblox', link: '/getting-started/welcome' },
          {
            text: 'Authorizing',
            items: [
              { text: 'Opencloud Authentication', link: '/getting-started/authorizing/opencloud-authentication' },
              // { text: 'Legacy  Authentication', link: '/getting-started/authorizing/legacy-authentication' },
            ]
          }
        ],
      },
      // {
      //   text: 'Guides',
      //   items: [
      //     {
      //       text: 'OpenCloud',
      //       items: [
      //       ]
      //     },
      //   ]
      // },
      {
        text: 'Packages',
        items: [
          { text: 'methodutil', link: '/getting-started/packages/methodutil' },
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/typical-developers/goblox' },
      { icon: 'discord', link: 'https://discord.gg/typical' },
      { icon: 'twitter', link: 'https://twitter.com/typicaldevelops' },
      { icon: 'roblox', link: 'https://www.roblox.com/communities/2700233/Typical-Developers' },
    ]
  }
})

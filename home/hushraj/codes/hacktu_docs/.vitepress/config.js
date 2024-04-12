import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "HackTU EMS",
  description: "HackTU Event management system",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    // nav: [
    //   { text: 'Home', link: '/' },
    //   { text: 'Examples', link: '/markdown-examples' }
    // ],

    sidebar: [
      {
        text: 'Links',
        items: [
          { text: 'Creative Computing Society', link: 'https://ccstiet.com' },
          { text: 'HackTU 5.0', link: 'https://helix.ccstiet.com/' },
          { text: 'Github Repo Link', link: 'https://github.com/saini128/eventAttendenceSystem' }
          
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/vuejs/vitepress' }
    ]
  }
})

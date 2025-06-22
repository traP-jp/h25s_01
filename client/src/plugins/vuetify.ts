/**
 * plugins/vuetify.ts
 *
 * Framework documentation: https://vuetifyjs.com`
 */

// Composables
import { createVuetify } from 'vuetify'
// Styles
import '@mdi/font/css/materialdesignicons.css'

import 'vuetify/styles'

// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
export default createVuetify({
  theme: {
    themes: {
      light: {
        dark: false,
        colors: {
          primary: '#FFC300',
          secondary: '#B2BDC7',
          warning: '#F26451',

          primaryText: '#EFF0F0',
          secondaryText: '#333E47',
          text: '#222425',

          background: '#FFFFFF',
        },
      },
    },
  },
})

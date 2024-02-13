/** @type {import('tailwindcss').Config} */
export default {
  content: ["./src/**/*.{html,js,jsx,tsx,ts}"],
  theme: {
    extend: {
      colors: {
        green: {
          '100': "#EAEFC5",
          '200': '#BDE09E',
          '300': '#90BB7C',
          '400': '#5F9256',
          '500': '#497042',
          '600': '#34502F',
          '700': '#1F301C',
        },
      },
    },
  },
  plugins: [],
}


import type { Config } from 'tailwindcss';

const config: Config = {
  content: [
    './pages/**/*.{js,ts,jsx,tsx,mdx}',
    './components/**/*.{js,ts,jsx,tsx,mdx}',
    './app/**/*.{js,ts,jsx,tsx,mdx}',
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          50: '#eef6ff',
          100: '#ddeaff',
          200: '#bcdbff',
          300: '#8cc3ff',
          400: '#5aa3ff',
          500: '#3385ff',
          600: '#006AFF',
          700: '#0056cc',
          800: '#004499',
          900: '#003366',
          950: '#001a33',
        },
        secondary: {
          50: '#e6fff9',
          100: '#ccfff4',
          200: '#99ffea',
          300: '#66ffe1',
          400: '#33ffd8',
          500: '#00C896',
          600: '#00a078',
          700: '#00785a',
          800: '#00583f',
          900: '#003d2e',
          950: '#001f18',
        },
      },
      borderRadius: {
        'xl': '12px',
        '2xl': '16px',
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
};

export default config;

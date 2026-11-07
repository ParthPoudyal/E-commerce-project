/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        studio: {
          bg: '#F9F9F6',          // Warm studio paper canvas
          surface: '#FFFFFF',     // Clean paper plate
          subtle: '#F2F2EC',      // Architectural warm grey fill
          border: '#E4E4DC',      // Hairline rule
          'border-dark': '#2B2B28', // Dark structural border
          ink: '#141413',         // Rich warm black
          graphite: '#4A4A45',    // Secondary text
          muted: '#84847D',       // Metadata & captions
          terracotta: '#B83A24',  // Live gavel accent / studio cinnabar
          'terracotta-light': '#F8EBE8',
          olive: '#2E5A3C',       // Reserve met / success stamp
          'olive-light': '#EBF2ED',
          khalti: '#5C2D91',      // Khalti Purple Stamp
          'khalti-light': '#F4EEFB',
        },
      },
      fontFamily: {
        serif: ['Newsreader', 'Playfair Display', 'Georgia', 'Cambria', 'serif'],
        sans: ['"Plus Jakarta Sans"', 'Inter', '-apple-system', 'sans-serif'],
        mono: ['"JetBrains Mono"', '"IBM Plex Mono"', '"Space Mono"', 'Menlo', 'Consolas', 'monospace'],
      },
      boxShadow: {
        'studio-sm': '0 1px 2px rgba(20, 20, 19, 0.04)',
        'studio-card': '0 2px 6px -1px rgba(20, 20, 19, 0.06), 0 1px 3px -1px rgba(20, 20, 19, 0.04)',
        'studio-lift': '0 12px 28px -4px rgba(20, 20, 19, 0.08), 0 4px 10px -2px rgba(20, 20, 19, 0.03)',
      },
      letterSpacing: {
        'studio-wide': '0.12em',
        'studio-wider': '0.18em',
      },
    },
  },
  plugins: [],
};

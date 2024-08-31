/** @type {import('tailwindcss').Config} */
export default {
  content: ["./src/**/*.{html,js,svelte,ts}"],
  theme: {
    extend: {
      colors: {
        cinema: {
          primary: '#000000', // Black
          secondary: '#1a1a1a', // Dark Gray
          highlight: '#B22222', // Deep Red
          text: '#FFFFFF', // White
          gold: '#FFD700', // Gold
          silver: '#C0C0C0', // Silver
          action: '#4169E1', // Royal Blue
          success: '#50C878', // Emerald Green
          warning: '#FF4500', // Sunset Orange
          goColor: '#007d9c',
        },
      },
    },
    plugins: [],
  },
}


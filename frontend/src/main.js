import htmx from 'htmx.org'

// Make htmx available globally
window.htmx = htmx

// Log when htmx is ready
document.addEventListener('DOMContentLoaded', () => {
  console.log('htmx loaded:', htmx.version)
})

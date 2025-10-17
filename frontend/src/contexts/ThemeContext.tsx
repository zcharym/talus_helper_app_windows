import React, { createContext, useContext, useEffect, useState } from 'react'
import { GetConfig, SaveConfig } from '@wailsjs/go/main/App'

type Theme = 'light' | 'dark' | 'auto'

interface ThemeContextType {
  theme: Theme
  isDark: boolean
  setTheme: (theme: Theme) => void
  toggleTheme: () => void
}

const ThemeContext = createContext<ThemeContextType | undefined>(undefined)

export const useTheme = () => {
  const context = useContext(ThemeContext)
  if (context === undefined) {
    throw new Error('useTheme must be used within a ThemeProvider')
  }
  return context
}

interface ThemeProviderProps {
  children: React.ReactNode
}

export const ThemeProvider: React.FC<ThemeProviderProps> = ({ children }) => {
  const [theme, setThemeState] = useState<Theme>('light')
  const [isDark, setIsDark] = useState(false)

  // Load theme from config on mount
  useEffect(() => {
    const loadTheme = async () => {
      try {
        const config = await GetConfig()
        setThemeState(config.Theme as Theme)
      } catch (error) {
        console.error('Failed to load theme config:', error)
        setThemeState('light')
      }
    }
    loadTheme()
  }, [])

  // Apply theme changes
  useEffect(() => {
    const applyTheme = async () => {
      let shouldBeDark = false

      if (theme === 'dark') {
        shouldBeDark = true
      } else if (theme === 'auto') {
        shouldBeDark = window.matchMedia('(prefers-color-scheme: dark)').matches
      }

      setIsDark(shouldBeDark)
      
      // Apply to document
      if (shouldBeDark) {
        document.documentElement.classList.add('dark')
      } else {
        document.documentElement.classList.remove('dark')
      }

      // Save to config
      try {
        const config = await GetConfig()
        const updatedConfig = { ...config, Theme: theme }
        await SaveConfig(updatedConfig)
      } catch (error) {
        console.error('Failed to save theme config:', error)
      }
    }

    applyTheme()
  }, [theme])

  // Listen for system theme changes when in auto mode
  useEffect(() => {
    if (theme !== 'auto') return

    const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)')
    const handleChange = () => {
      const shouldBeDark = mediaQuery.matches
      setIsDark(shouldBeDark)
      
      if (shouldBeDark) {
        document.documentElement.classList.add('dark')
      } else {
        document.documentElement.classList.remove('dark')
      }
    }

    mediaQuery.addEventListener('change', handleChange)
    return () => mediaQuery.removeEventListener('change', handleChange)
  }, [theme])

  const setTheme = (newTheme: Theme) => {
    setThemeState(newTheme)
  }

  const toggleTheme = () => {
    setThemeState(prev => prev === 'light' ? 'dark' : 'light')
  }

  const value: ThemeContextType = {
    theme,
    isDark,
    setTheme,
    toggleTheme
  }

  return (
    <ThemeContext.Provider value={value}>
      {children}
    </ThemeContext.Provider>
  )
}

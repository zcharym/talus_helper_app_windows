import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import TodoList from './components/TodoList'
import Settings from './components/Settings'
import Tabs, { navigationTabs } from './components/Tabs'
import { ThemeProvider } from './contexts/ThemeContext'

function App() {
  return (
    <ThemeProvider>
      <Router>
        <div className="min-h-screen bg-gray-50 dark:bg-gray-900 transition-colors duration-200">
          <Tabs tabs={navigationTabs} />
          <main className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
            <Routes>
              <Route path="/" element={<TodoList />} />
              <Route path="/settings/*" element={<Settings />} />
            </Routes>
          </main>
        </div>
      </Router>
    </ThemeProvider>
  )
}

export default App

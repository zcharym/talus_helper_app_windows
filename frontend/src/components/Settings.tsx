import { Routes, Route, Navigate } from 'react-router-dom'
import SettingsSidebar from './settings/SettingsSidebar'
import AppearanceSettings from './settings/AppearanceSettings'
import TodoSettings from './settings/TodoSettings'
import OCRSettings from './settings/OCRSettings'
import GeneralSettings from './settings/GeneralSettings'
import LanguageSettings from './settings/LanguageSettings'

function Settings() {
  return (
    <div className="flex h-full">
      {/* Sidebar */}
      <SettingsSidebar />
      
      {/* Main Content */}
      <div className="flex-1 overflow-auto">
        <div className="p-6">
          <Routes>
            <Route path="/" element={<Navigate to="/settings/appearance" replace />} />
            <Route path="/appearance" element={<AppearanceSettings />} />
            <Route path="/todos" element={<TodoSettings />} />
            <Route path="/ocr" element={<OCRSettings />} />
            <Route path="/general" element={<GeneralSettings />} />
            <Route path="/language" element={<LanguageSettings />} />
          </Routes>
        </div>
      </div>
    </div>
  )
}

export default Settings

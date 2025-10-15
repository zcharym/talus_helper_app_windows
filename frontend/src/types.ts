// Temporary types until Wails bindings are properly generated
export interface Todo {
  id: string
  text: string
  completed: boolean
  createdAt: string
}

export interface AppConfig {
  theme: string
  autoSave: boolean
  notifications: boolean
  openAIAPIKey: string
  openAIBaseURL: string
  defaultTodoCategory: string
  maxTodos: number
  language: string
}

// Mock functions for development
export const GetTodos = async (): Promise<Todo[]> => {
  // Mock implementation
  return []
}

export const AddTodo = async (text: string): Promise<Todo> => {
  // Mock implementation
  return {
    id: Math.random().toString(36),
    text,
    completed: false,
    createdAt: new Date().toISOString()
  }
}

export const UpdateTodo = async (id: string, text: string, completed: boolean): Promise<Todo> => {
  // Mock implementation
  return {
    id,
    text,
    completed,
    createdAt: new Date().toISOString()
  }
}

export const DeleteTodo = async (_id: string): Promise<void> => {
  // Mock implementation
}

export const GetConfig = async (): Promise<AppConfig> => {
  // Mock implementation
  return {
    theme: 'light',
    autoSave: true,
    notifications: true,
    openAIAPIKey: '',
    openAIBaseURL: 'https://api.moonshot.cn/v1',
    defaultTodoCategory: 'General',
    maxTodos: 100,
    language: 'en'
  }
}

export const SaveConfig = async (_config: AppConfig): Promise<void> => {
  // Mock implementation
}

export const OCRFromClipboard = async (): Promise<string> => {
  // Mock implementation
  return "Mock OCR result"
}

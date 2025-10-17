import { useState, useEffect } from 'react'
import { Plus, Edit2, Trash2, Check, X, Clipboard } from 'lucide-react'
import { GetTodos, AddTodo, UpdateTodo, DeleteTodo, OCRFromClipboard } from '@wailsjs/go/main/App'
import { Todo } from '../types'

function TodoList() {
  const [todos, setTodos] = useState<Todo[]>([])
  const [newTodoText, setNewTodoText] = useState('')
  const [editingId, setEditingId] = useState<string | null>(null)
  const [editingText, setEditingText] = useState('')
  const [loading, setLoading] = useState(true)
  const [ocrLoading, setOcrLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  // Load todos on component mount
  useEffect(() => {
    loadTodos()
  }, [])

  const loadTodos = async () => {
    try {
      setLoading(true)
      const todosData = await GetTodos()
      setTodos(todosData || [])
    } catch (error) {
      console.error('Failed to load todos:', error)
    } finally {
      setLoading(false)
    }
  }

  const handleAddTodo = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!newTodoText.trim()) return

    try {
      const newTodo = await AddTodo(newTodoText.trim())
      setTodos(prev => [...prev, newTodo])
      setNewTodoText('')
    } catch (error) {
      console.error('Failed to add todo:', error)
    }
  }

  const handleToggleTodo = async (id: string) => {
    try {
      const todo = todos.find(t => t.id === id)
      if (!todo) return

      const updatedTodo = await UpdateTodo(id, todo.text, !todo.completed)
      setTodos(prev => prev.map(t => t.id === id ? updatedTodo : t))
    } catch (error) {
      console.error('Failed to toggle todo:', error)
    }
  }

  const handleEditTodo = (id: string, text: string) => {
    setEditingId(id)
    setEditingText(text)
  }

  const handleSaveEdit = async (id: string) => {
    if (!editingText.trim()) return

    try {
      const todo = todos.find(t => t.id === id)
      if (!todo) return

      const updatedTodo = await UpdateTodo(id, editingText.trim(), todo.completed)
      setTodos(prev => prev.map(t => t.id === id ? updatedTodo : t))
      setEditingId(null)
      setEditingText('')
    } catch (error) {
      console.error('Failed to update todo:', error)
    }
  }

  const handleCancelEdit = () => {
    setEditingId(null)
    setEditingText('')
  }

  const handleDeleteTodo = async (id: string) => {
    try {
      await DeleteTodo(id)
      setTodos(prev => prev.filter(t => t.id !== id))
    } catch (error) {
      console.error('Failed to delete todo:', error)
    }
  }

  const handleOCRFromClipboard = async () => {
    try {
      setOcrLoading(true)
      setError(null)
      
      const extractedText = await OCRFromClipboard()
      setNewTodoText(extractedText)
    } catch (error) {
      console.error('OCR failed:', error)
      setError(error instanceof Error ? error.message : 'Failed to extract text from clipboard')
    } finally {
      setOcrLoading(false)
    }
  }

  const completedCount = todos.filter(todo => todo.completed).length
  const totalCount = todos.length

  if (loading) {
    return (
      <div className="flex justify-center items-center h-64">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary-600"></div>
      </div>
    )
  }

  return (
    <div className="px-4 py-6 sm:px-0">
      <div className="max-w-2xl mx-auto">
        <div className="mb-8">
          <h2 className="text-3xl font-bold text-gray-900 mb-2">Todo List</h2>
          <p className="text-gray-600">
            {totalCount === 0 
              ? "No tasks yet. Add one below to get started!" 
              : `${completedCount} of ${totalCount} tasks completed`
            }
          </p>
        </div>

        {/* Error Message */}
        {error && (
          <div className="mb-4 p-4 bg-red-50 border border-red-200 text-red-800 rounded-lg">
            {error}
          </div>
        )}

        {/* Add Todo Form */}
        <form onSubmit={handleAddTodo} className="mb-8">
          <div className="flex gap-2">
            <input
              type="text"
              value={newTodoText}
              onChange={(e) => setNewTodoText(e.target.value)}
              placeholder="What needs to be done?"
              className="input-field flex-1"
              maxLength={200}
            />
            <button
              type="button"
              onClick={handleOCRFromClipboard}
              disabled={ocrLoading}
              className="btn-secondary flex items-center gap-2"
              title="Read text from clipboard image"
            >
              {ocrLoading ? (
                <div className="w-4 h-4 border-2 border-gray-300 border-t-gray-600 rounded-full animate-spin" />
              ) : (
                <Clipboard className="w-4 h-4" />
              )}
              {ocrLoading ? 'Reading...' : 'Read from Clipboard'}
            </button>
            <button
              type="submit"
              className="btn-primary flex items-center gap-2"
              disabled={!newTodoText.trim()}
            >
              <Plus className="w-4 h-4" />
              Add
            </button>
          </div>
        </form>

        {/* Todo List */}
        <div className="space-y-2">
          {todos.length === 0 ? (
            <div className="text-center py-12">
              <div className="text-gray-400 mb-4">
                <Check className="w-16 h-16 mx-auto" />
              </div>
              <h3 className="text-lg font-medium text-gray-900 mb-2">All done!</h3>
              <p className="text-gray-500">Add a task above to get started.</p>
            </div>
          ) : (
            todos.map((todo) => (
              <div
                key={todo.id}
                className={`card flex items-center gap-3 ${
                  todo.completed ? 'opacity-75' : ''
                }`}
              >
                <button
                  onClick={() => handleToggleTodo(todo.id)}
                  className={`flex-shrink-0 w-5 h-5 rounded border-2 flex items-center justify-center transition-colors ${
                    todo.completed
                      ? 'bg-primary-600 border-primary-600 text-white'
                      : 'border-gray-300 hover:border-primary-500'
                  }`}
                >
                  {todo.completed && <Check className="w-3 h-3" />}
                </button>

                {editingId === todo.id ? (
                  <div className="flex-1 flex gap-2">
                    <input
                      type="text"
                      value={editingText}
                      onChange={(e) => setEditingText(e.target.value)}
                      className="input-field flex-1"
                      autoFocus
                      onKeyDown={(e) => {
                        if (e.key === 'Enter') handleSaveEdit(todo.id)
                        if (e.key === 'Escape') handleCancelEdit()
                      }}
                    />
                    <button
                      onClick={() => handleSaveEdit(todo.id)}
                      className="btn-primary"
                      disabled={!editingText.trim()}
                    >
                      <Check className="w-4 h-4" />
                    </button>
                    <button
                      onClick={handleCancelEdit}
                      className="btn-secondary"
                    >
                      <X className="w-4 h-4" />
                    </button>
                  </div>
                ) : (
                  <>
                    <span
                      className={`flex-1 ${
                        todo.completed ? 'line-through text-gray-500' : 'text-gray-900'
                      }`}
                    >
                      {todo.text}
                    </span>
                    <div className="flex gap-1">
                      <button
                        onClick={() => handleEditTodo(todo.id, todo.text)}
                        className="p-2 text-gray-400 hover:text-primary-600 transition-colors"
                        title="Edit"
                      >
                        <Edit2 className="w-4 h-4" />
                      </button>
                      <button
                        onClick={() => handleDeleteTodo(todo.id)}
                        className="p-2 text-gray-400 hover:text-red-600 transition-colors"
                        title="Delete"
                      >
                        <Trash2 className="w-4 h-4" />
                      </button>
                    </div>
                  </>
                )}
              </div>
            ))
          )}
        </div>
      </div>
    </div>
  )
}

export default TodoList

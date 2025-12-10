import { useEffect, useMemo, useState } from "react";
import * as AppAPI from "../../wailsjs/go/main/App";

interface ScreenshotFile {
  name: string;
  path: string;
  size: number;
  modTime: string | Date;
}

function formatBytes(bytes: number): string {
  const sizes = ["B", "KB", "MB", "GB"];
  if (bytes === 0) return "0 B";
  const i = Math.floor(Math.log(bytes) / Math.log(1024));
  const value = bytes / Math.pow(1024, i);
  return `${value.toFixed(1)} ${sizes[i]}`;
}

export default function Screenshots() {
  const [items, setItems] = useState<ScreenshotFile[]>([]);
  const [selected, setSelected] = useState<ScreenshotFile | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const previewSrc = useMemo(() => {
    if (!selected) return "";
    // Use file path as src; Wails loads via OS file protocol
    return `file://${selected.path}`;
  }, [selected]);

  async function load() {
    setLoading(true);
    setError(null);
    try {
      const list = await AppAPI.ListScreenshots();
      setItems(list);
      setSelected(list[0] ?? null);
    } catch (e: any) {
      setError(e?.message || "Failed to load screenshots");
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    load();
  }, []);

  return (
    <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
      <div className="md:col-span-1 border rounded-lg overflow-hidden bg-white dark:bg-gray-800">
        <div className="p-3 border-b flex items-center justify-between">
          <h2 className="font-medium">Screenshots</h2>
          <button
            onClick={load}
            className="text-sm px-2 py-1 rounded bg-gray-100 dark:bg-gray-700 hover:bg-gray-200 dark:hover:bg-gray-600"
          >
            Refresh
          </button>
        </div>
        <div className="divide-y max-h-[70vh] overflow-auto">
          {loading && <div className="p-4 text-sm text-gray-500">Loading…</div>}
          {error && <div className="p-4 text-sm text-red-600">{error}</div>}
          {!loading && items.length === 0 && !error && (
            <div className="p-4 text-sm text-gray-500">
              No screenshots found.
            </div>
          )}
          {items.map((it) => {
            const isActive = selected?.path === it.path;
            return (
              <button
                key={it.path}
                onClick={() => setSelected(it)}
                className={`w-full text-left px-3 py-2 focus:outline-none transition-colors ${
                  isActive
                    ? "bg-primary-50 dark:bg-primary-900/30"
                    : "hover:bg-gray-50 dark:hover:bg-gray-700"
                }`}
              >
                <div className="flex items-center justify-between">
                  <div className="truncate">
                    <div className="truncate text-sm font-medium">
                      {it.name}
                    </div>
                    <div className="text-xs text-gray-500">
                      {formatBytes(it.size)} ·{" "}
                      {new Date(it.modTime).toLocaleString()}
                    </div>
                  </div>
                </div>
              </button>
            );
          })}
        </div>
      </div>

      <div className="md:col-span-2 border rounded-lg overflow-hidden bg-white dark:bg-gray-800">
        <div className="p-3 border-b flex items-center justify-between">
          <h2 className="font-medium">Preview</h2>
          <div className="flex gap-2">
            <button
              disabled={!selected}
              onClick={() => selected && AppAPI.OpenScreenshot(selected.path)}
              className="text-sm px-2 py-1 rounded bg-gray-100 dark:bg-gray-700 hover:bg-gray-200 dark:hover:bg-gray-600 disabled:opacity-50"
            >
              Open
            </button>
            <button
              disabled
              className="text-sm px-2 py-1 rounded bg-gray-100 dark:bg-gray-700 opacity-60"
            >
              Reveal
            </button>
            <button
              disabled
              className="text-sm px-2 py-1 rounded bg-red-100 dark:bg-red-900/30 text-red-700 dark:text-red-300 opacity-60"
            >
              Delete
            </button>
          </div>
        </div>
        <div className="p-3 flex items-center justify-center min-h-[60vh]">
          {selected ? (
            // eslint-disable-next-line @next/next/no-img-element
            <img
              alt={selected.name}
              src={previewSrc}
              className="max-h-[70vh] max-w-full object-contain rounded"
            />
          ) : (
            <div className="text-sm text-gray-500">
              Select a screenshot to preview
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

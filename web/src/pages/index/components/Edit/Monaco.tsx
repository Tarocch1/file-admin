import { useState, useEffect, useRef } from 'react'
import * as monaco from 'monaco-editor'

import { getLanguage } from './utils/getLanguage'

export interface IMonacoProp {
  value: string
  target: string
  onChange: (value: string) => void
}

export default function Monaco({ value, target, onChange }: IMonacoProp) {
  const [editor, setEditor] =
    useState<monaco.editor.IStandaloneCodeEditor | null>(null)
  const monacoEl = useRef<HTMLDivElement>(null)

  useEffect(() => {
    if (!editor) {
      const model = monaco.editor.createModel(value, getLanguage(target))
      model.setEOL(monaco.editor.EndOfLineSequence.LF)
      const _editor = monaco.editor.create(monacoEl.current!, {
        fontSize: 14,
        tabSize: 2,
        roundedSelection: false,
        emptySelectionClipboard: false,
        model,
      })
      _editor.onDidChangeModelContent((e) => {
        onChange(_editor.getValue())
      })
      setEditor(_editor)
    }

    return () => editor?.dispose()
  }, [])

  return <div style={{ height: '70vh' }} ref={monacoEl}></div>
}

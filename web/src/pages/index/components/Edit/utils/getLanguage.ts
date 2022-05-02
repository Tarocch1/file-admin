const languageMap = {
  bat: ['bat'],
  cpp: ['c', 'cpp', 'cc', 'cxx', 'c++', 'h', 'hpp', 'hxx'],
  css: ['css'],
  dart: ['dart'],
  dockerfile: ['dockerfile'],
  go: ['go'],
  html: ['html'],
  ini: ['ini'],
  java: ['java'],
  javascript: ['js', 'jsx', 'javascript'],
  json: ['json', 'jsonc'],
  kotlin: ['kt'],
  less: ['less'],
  lua: ['lua'],
  markdown: ['md', 'markdown'],
  perl: ['pl'],
  php: ['php'],
  powershell: ['ps1'],
  protobuf: ['pb'],
  python: ['py', 'python'],
  ruby: ['rb'],
  rust: ['rs'],
  scala: ['scala'],
  scss: ['scss'],
  shell: ['sh', 'zsh', 'bash'],
  sql: ['sql'],
  swift: ['swift'],
  typescript: ['ts', 'tsx', 'typescript'],
  xml: ['xml'],
  yaml: ['yaml', 'yml'],
}

export function getLanguage(target: string) {
  const ext = target.split('.').pop()
  if (ext) {
    const lower = ext.toLowerCase()
    for (const language of Object.keys(languageMap)) {
      if (languageMap[language as keyof typeof languageMap].includes(lower)) {
        return language
      }
    }
  }
  return undefined
}

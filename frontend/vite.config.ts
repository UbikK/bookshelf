import {defineConfig} from 'vite'
import {svelte} from '@sveltejs/vite-plugin-svelte'
import {resolve} from 'path'
// https://vitejs.dev/config/
const projectRootDir = resolve(__dirname);
export default defineConfig({
  plugins: [svelte()],
  resolve: {alias:[
    {
      find: '@',
      replacement: resolve(projectRootDir)
    },
    {
      find: 'wails',
      replacement: resolve(projectRootDir, 'wailsjs', 'go')
    }
  ]} ,server: {
    watch: {
      ignored: ['**/database.db'],
    },
  }
})

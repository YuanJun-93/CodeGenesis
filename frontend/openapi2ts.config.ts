import { resolve } from 'path';

export default {
  requestLibPath: "import request from '@/request'",
  schemaPath: resolve(__dirname, '../api/doc/swagger.json'),
  serversPath: './src',
}

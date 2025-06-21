import { HttpResponse } from 'msw'
import { createOpenApiHttp } from 'openapi-msw'
import type { paths } from '@/api/schema'

export const http = createOpenApiHttp<paths>()
export const handlers = [
  http.get('/events/{eventID}', () => {
    return HttpResponse.json({
      title: '春ハッカソン2025',
      description: '参加登録してね！！',
      date: '2025-06-21',
      is_open: true,
      is_me_attendee: true,
      admins: ['cp20','ogu_kazemiya'],
      attendees: ['cp20','ogu_kazemiya','Eraxyso', 'hachimitsu', 'ten_ten', 'miyamon', 'Naph1'],
      event_id: '3fa85f64-5717-4562-b3fc-2c963f66afa6',
      is_deleted: false,
      updated_at: '2025-06-21T07:13:04.294Z',
      created_at: '2025-06-21T07:13:04.294Z',
    })
  }),
]

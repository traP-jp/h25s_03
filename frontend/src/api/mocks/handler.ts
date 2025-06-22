import { HttpResponse } from 'msw'
import { createOpenApiHttp } from 'openapi-msw'
import type { paths } from '@/api/schema'

export const http = createOpenApiHttp<paths>()
export const handlers = [
  http.get('/events', () => {
    return HttpResponse.json([
      {
        title: 'Event 1',
        description: 'Description for Event 1',
        date: '2023-10-01',
        is_open: true,
        is_me_admin: false,
      },
    ])
  }),
  http.get('/events/{eventID}/lotteries/{lotteryID}', () => {
    return HttpResponse.json({
      lottery_id: '3fa85f64-5717-4562-b3fc-2c963f66afa6',
      event_id: '3fa85f64-5717-4562-b3fc-2c963f66afa6',
      title: '焼肉食べ放題',
      is_deleted: false,
      created_at: '2025-06-21T06:06:28.642Z',
      updated_at: '2025-06-21T06:06:28.642Z',
      winners: ['miyamon', 'ogu_kazemiya', 'ten_ten'],
    })
  }),
]

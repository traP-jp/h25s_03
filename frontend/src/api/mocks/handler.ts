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
        date: '2026-10-01',
        is_open: false,
        is_me_attendee: true,
        admins: ['ten_ten']
      },
      {
        title: 'Event 2',
        description: 'Description for Event 2',
        date: '2023-10-01',
        is_open: true,
        is_me_attendee: true,
        admins: ['ten_ten']
      }
    ])
  }),
  // http.post('/events', () => {
  //   return HttpResponse.json([
  //     {
  //       event_id: '1234',
  //       title: 'Event 1',
  //       description: 'Description for Event 1',
  //       date: '2023-10-01',
  //       is_open: true,
  //       is_me_admin: false,
  //     }
  //   ])
  // })
]

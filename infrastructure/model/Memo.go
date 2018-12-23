package memo

type Memo struct{
  Id string `json:"id"`
  Title string  `json:"title"`
  Text string  `json:"text"`
  Flag bool    `json:"check"`
  Date string    `json:"date"`
}

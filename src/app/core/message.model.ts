export class Message {
  constructor(
    public id: string,
    public date: Date,
    public subject: string,
    public author: string,
  ) {}
}

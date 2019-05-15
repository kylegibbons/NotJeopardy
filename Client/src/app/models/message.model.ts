export interface Message {
    messageType: string;
    gameId: string;
    timestamp: Date;
    payload: any;
}

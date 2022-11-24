import { Client } from "@xmtp/xmtp-js";
import web3 from "web3"


function addressChecksum(recipientAddress: any) {
  return web3.utils.toChecksumAddress(recipientAddress);
}

export default class XMTPManager {
  static clientInstance: any = null;

  static connected = () => this.clientInstance !== null;


  static async getMessages(recipientAddress: any) {
    if (this.connected()) {
      const conversation = await this.clientInstance.conversations.newConversation(recipientAddress)
      const messages = await conversation.messages();
      return messages
    }
    throw new Error("XMTP not connected!");
  }

  static async sendMessage(recipientAddress: any, message: any) {
    if (this.connected()) {
      const conversation =
        await this.clientInstance.conversations.newConversation(addressChecksum(recipientAddress));
      await conversation.send(message);
      return;
    }
    throw new Error("XMTP not connected!");
  }

  static async getConversations() {
    if (this.connected()) {
      console.log(this.clientInstance);
      return await this.clientInstance.conversations.list();
    }
    throw new Error("XMTP not connected!");
  }

  static async getInstance(wallet: any) {
    if (XMTPManager.clientInstance === null) {
      XMTPManager.clientInstance = await Client.create(wallet);
    }
    return this.clientInstance;
  }
}

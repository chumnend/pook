function toSentenceCase(message: string): string {
    if (!message) return '';
    return message.charAt(0).toUpperCase() + message.slice(1).toLowerCase();
};
  
export default toSentenceCase;

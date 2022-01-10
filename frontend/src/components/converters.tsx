// FileオブジェクトをBase64のstringに変換
const toBase64 = async (file: File) => {
    return fileToBase64(file)
           .then((result: any) => {
      return result;
    });
}

const fileToBase64 = async (file: File) => {
    return new Promise(resolve => {
      const reader = new FileReader();
  
      // Read file content on file loaded event
      reader.onload = (event: any) => {
        resolve(event.target.result);
      };
      
      // Convert data to base64 
      reader.readAsDataURL(file);
    });
};

const createJpegFile4Base64 = function (base64: string, name: string) {
  // base64のデコード
  const bin = atob(base64.replace(/^.*,/, ''));
  // バイナリデータ化
  const buffer = new Uint8Array(bin.length);
  for (let i = 0; i < bin.length; i++) {
      buffer[i] = bin.charCodeAt(i);
  }
  // ファイルオブジェクト生成(この例ではjpegファイル)
  return new File([buffer.buffer], name, {type: "image/jpg"});
};

export { toBase64, createJpegFile4Base64 }
import axios from "axios";
import React, { useState, ChangeEvent } from "react";

const PostImage = () => {
    const [submitData, setSubmitData] = useState<string>("")
    const [imgURL, setImgURL]         = useState<string>("")
    const [replyImgURL, setReplyURL]  = useState<string>("")

    const handleImageInput = async (e: ChangeEvent<HTMLInputElement>) => {
        // 入力用の画像が変更されるたびに、表示用の画像URLを更新し　提出用データも更新
        if (!e.target.files) return;
        const _rawdata: File = e.target.files[0];
        setImgURL( URL.createObjectURL(_rawdata) );
        setReplyURL("");

        // 画像をBase64エンコードしてsubmitDataにセットする
        if (!_rawdata) return;
        const base64file  = await toBase64(_rawdata)
        //提出用データのBodyの中身を作成
        const _submitData = JSON.stringify( {
                                text: base64file
                            } ) 
        setSubmitData( _submitData )
    }

    const handleSubmitData = async () => {
        if (!submitData) return;
        axios({
            method: "post",
            url:    "sample",
            data:    submitData
        })
        .catch(results => {
            alert("an error occured")
        });
    }

    return (
        <div className="container">
            <div className="title">画像</div>
            <input  type="file"
                    name="example"
                    accept="image/*"
                    onChange={handleImageInput}
            />
            
            <div className="images">
                <img src={ imgURL } alt="" />
                <img src={ replyImgURL } alt="" />
            </div>

            <div className="buttons">
                <button type="button"
                        onClick={handleSubmitData}
                >submit</button>
            </div>
        </div>
    )
}

export default PostImage


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

//const createJpegFile4Base64 = function (base64: string, name: string) {
//  // base64のデコード
//  const bin = atob(base64.replace(/^.*,/, ''));
//  // バイナリデータ化
//  const buffer = new Uint8Array(bin.length);
//  for (let i = 0; i < bin.length; i++) {
//      buffer[i] = bin.charCodeAt(i);
//  }
//  // ファイルオブジェクト生成(この例ではjpegファイル)
//  return new File([buffer.buffer], name, {type: "image/jpg"});
//};
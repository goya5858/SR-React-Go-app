import axios from "axios";
import React, { useState, ChangeEvent } from "react";
import { toBase64, createJpegFile4Base64 } from "./converters";

const PostImage = () => {
    const [submitData, setSubmitData] = useState<string>("")
    const [imgURL, setImgURL]         = useState<string>("")
    const [replyImgURL, setReplyImgURL]  = useState<string>("")

    const handleImageInput = async (e: ChangeEvent<HTMLInputElement>) => {
        // 入力用の画像が変更されるたびに、表示用の画像URLを更新し　提出用データも更新
        if (!e.target.files) return;
        const _rawdata: File = e.target.files[0];
        setImgURL( URL.createObjectURL(_rawdata) );
        setReplyImgURL("");

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
            url:    "http://localhost/api/submit",
            data:    submitData
        })
        .then(res => {
            console.log("response: ", res);
            let new_Img:File      = createJpegFile4Base64( res.data, "new_img" )
            setReplyImgURL( URL.createObjectURL(new_Img) )
        })
        .catch(results => {
            alert(results)
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
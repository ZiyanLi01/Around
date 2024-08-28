import React, { useState } from 'react';
import axios from 'axios';
import { Modal, Button, Input } from 'antd';
import { GENERATE_IMAGE_URL, TOKEN_KEY } from '../constants';

function ImageGenerator() {
    const [description, setDescription] = useState('');
    const [imageUrl, setImageUrl] = useState('');
    const [loading, setLoading] = useState(false);
    const [isModalVisible, setIsModalVisible] = useState(false);

    const showModal = () => {
        setIsModalVisible(true);
    };

    const handleOk = async () => {
        setLoading(true);

        const token = localStorage.getItem(TOKEN_KEY); // Retrieve the token from local storage

        try {
            const response = await axios.post(
                GENERATE_IMAGE_URL,
                { description: description },
                {
                    headers: {
                        Authorization: `Bearer ${token}`, // Include the token in the Authorization header
                    },
                }
            );
            console.log("Full Response:", response.data);
            console.log("Image URL:", response.data.imageUrl); // Ensure you're accessing the correct property from the backend response
            // static test ImageUrl
            //setImageUrl('https://oaidalleapiprodscus.blob.core.windows.net/private/org-OfevZPM1cCub57CTJ3CEje8d/user-5AdjYI9XsOtfbDCm74Xu0I9f/img-Dc2WZl1l47z9hSH0DUpGpmPd.png?st=2024-08-27T22%3A12%3A39Z&se=2024-08-28T00%3A12%3A39Z&sp=r&sv=2024-08-04&sr=b&rscd=inline&rsct=image/png&skoid=d505667d-d6c1-4a0a-bac7-5c84a87759f8&sktid=a48cca56-e6da-484e-a814-9c849652bcb3&skt=2024-08-27T23%3A12%3A39Z&ske=2024-08-28T23%3A12%3A39Z&sks=b&skv=2024-08-04&sig=e/UdA8wC1NfTomIe4EZSylC5zrxWlsbn%2BLZTvCaRHrY%3D');
            setImageUrl(response.data.imageUrl); // Update the state with the image URL returned from the backend
        } catch (error) {
            console.error('Error generating image:', error);
        }

        setLoading(false);
    };

    const handleCancel = () => {
        setIsModalVisible(false);
        setDescription('');
        setImageUrl('');
    };

    return (
        <>
            <Button type="primary" onClick={showModal}>
                AI Generate Post
            </Button>
            <Modal
                title="AI Image Generator"
                visible={isModalVisible}
                onOk={handleOk}
                onCancel={handleCancel}
                confirmLoading={loading}
                okText={loading ? 'Generating...' : 'Generate Image'}
            >
                <Input 
                    value={description} 
                    onChange={(e) => setDescription(e.target.value)} 
                    placeholder="Enter image description" 
                    disabled={loading}
                />
                {imageUrl && (
                    <div style={{ marginTop: '20px', textAlign: 'center' }}>
                        <h3>Generated Image:</h3>
                        <img 
                            src={imageUrl} 
                            alt="Generated" 
                            style={{ maxWidth: '100%', height: 'auto' }} 
                        />
                        <a href={imageUrl} download style={{ display: 'block', marginTop: '10px' }}>
                            Download Image
                        </a>
                    </div>
                )}
            </Modal>
        </>
    );
}

export default ImageGenerator;

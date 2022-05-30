package graphics

var defaultVert2d = f32Bytes(
	0, 0, 0, 0, // top left
	+1, 0, 1, 0, // top right
	0, -1, 0, 1, // bottom left
	+1, -1, 1, 1, // bottom right
)

var vs2d = `#version 100

#define PI 3.14159265359

uniform mat3 mvp;
uniform mat3 uvp;

attribute  vec2 vert;
attribute vec2 clr_vert;

varying vec2 Tex2DCoords;
varying float fillColor;
varying vec4 color;


mat4 scale(float x, float y, float z) {
	return mat4(
	x, 0, 0, 0, 
	0, y, 0, 0, 
	0, 0, z, 0, 
	0, 0, 0, 1);
}

mat4 translate(float x, float y, float z){
	return mat4(
	1, 0, 0, 0, 
	0, 1, 0, 0, 
	0, 0, 1, 0, 
	x, y, z, 1);
}

mat4 ortho(float left, float right, float bottom, float top, float near, float far)  {
	float rml = right - left; 
	float tmb = top - bottom;
	float fmn = far - near;
	return mat4(
		2.0 / rml, 0, 0, 0,
		0, 2.0 / tmb, 0, 0, 
		0, 0, -2.0 / fmn, 0,
		-(right + left) / rml, -(top + bottom) / tmb, -(far + near) / fmn, 1
	);
}

mat4 lookat(vec3 eye, vec3 center, vec3 up) {

	vec3 d = normalize(eye-center);
	vec3 r = normalize(cross(up, d));
	vec3 u = cross(d, r);

	mat4 m = mat4(
		r[0], u[0], -d[0], 0,
		r[1], u[1], -d[1], 0,
		r[2], u[2], -d[2], 0,
		0, 0, 0, 1
	);

	return translate(-eye[0], -eye[1], -eye[2]) * m;
}

mat2 scale2d(vec2 _scale){
	return mat2(_scale.x,0.0,0.0, _scale.y);
}

mat4 rotate2d(float _angle){
	_angle = (_angle * PI) / 180.0;
    return mat4(cos(_angle), -sin(_angle), 0, 0,
	                sin(_angle), cos(_angle), 0 ,0,
					0,0,1,0,
					0,0,0,1);
}


void main(){
	
	
	float vw = mvp[0].x;
	float vh = mvp[1].x;
	float w = mvp[0].y; 
	float h = mvp[1].y; 
	float px = mvp[0].z;
	float py = mvp[1].z;
	float sx = mvp[2].x;
	float sy = mvp[2].y;
	float angle = mvp[2].z;

	float tw = uvp[0].x; 
	float th = uvp[0].y; 
	float minx = uvp[1].x; 
	float miny = uvp[1].y; 

	float upx = (minx)/tw;
	float upy = (miny)/tw;       

	vec2 clr = clr_vert * scale2d(vec2(w/tw, h/th));
	Tex2DCoords = vec2(clr.x + upx, clr.y + upy);

	fillColor =uvp[0].z;

	color = vec4(uvp[1].z, uvp[2]);


	// Rectangle 
	mat2 size = mat2(w/vw , 0, 0, h/vh);;
	gl_Position = vec4(vert *size, 0, 1);
	gl_Position = rotate2d(angle) * scale(sx, sy, 1.0) * gl_Position;
	gl_Position = translate(px/vw, -py/vh, 0.0) * gl_Position;
	gl_Position = ortho(0.0, 1.0, -1.0, 0.0, -1.0, 0.1) * gl_Position;
}

`

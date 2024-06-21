package cn.net.bhe.eciportal.service;

import cn.net.bhe.eciportal.config.ResProp;
import cn.net.bhe.eciportal.model.dto.EciDTO;
import cn.net.bhe.eciportal.model.vo.EciVO;
import cn.net.bhe.mutil.HttpUtils;
import com.alibaba.fastjson2.JSON;
import org.springframework.stereotype.Service;
import org.springframework.web.util.UriComponentsBuilder;

import java.net.URI;
import java.util.List;

@Service
public class EciServiceImpl implements EciService {

    private final ResProp resProp;

    public EciServiceImpl(ResProp resProp) {
        this.resProp = resProp;
    }

    @Override
    public List<EciVO> queryEciList() {
        String str = HttpUtils.get(resProp.getApiEciList());
        return JSON.parseArray(str, EciVO.class);
    }

    @Override
    public void addEci(EciDTO eciDTO) {
        HttpUtils.post(resProp.getApiEci(), JSON.toJSONString(eciDTO));
    }

    @Override
    public void deleteEci(String name) {
        URI uri = UriComponentsBuilder.fromUriString(resProp.getApiEci())
                .queryParam("name", name)
                .build()
                .toUri();
        HttpUtils.delete(uri.toString());
    }

}
